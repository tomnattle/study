#include <signal.h>
#include <stdio.h>
#include <string.h>
#include <sys/types.h>
#include <unistd.h>
#include <sys/wait.h>

//全局变量
sig_atomic_t sigusr1_count = 0;
sig_atomic_t child_exit_status;

void handler(int sig_num){
    ++sigusr1_count;
}

typedef int pid_t; 

void clean_up_child_process() {
    int status;
    wait(&status);
    child_exit_status = status;
}

void testSignal(){
    struct sigaction sa;

    memset (&sa, 0, sizeof(sa));
    sa.sa_handler = &handler;
    sigaction (SIGUSR1, &sa, NULL);

    printf ("SIGUSR1 was raised %d times\n", sigusr1_count);


    struct sigaction sigchld_action;
    memset (&sigchld_action, 0, sizeof (sigchld_action));
    sigchld_action.sa_handler = &clean_up_child_process;
    // 注册子进程退出⌚️ 避免wait函数 block主进程的执行
    sigaction (SIGCHLD, &sigchld_action, NULL);

}

void testWait(){
    int child_status;
    char* args[] = {
        "ls",
        "-l",
        ".",
        "NULL",
    };

    pid_t parent_pid;
    pid_t child_pid;
    parent_pid = getpid();
    child_pid = fork();

    if (child_pid != 0){
        //父进程
        printf("p:%d\n", parent_pid);
        printf("p:child id is %d\n", child_pid);
    }else{
        //子进程
        printf("c:%d\n", getpid());
        execvp("ls", args);
    }
    /**
    //wait3
    wait(&child_status);
    if (WIFEXITED(child_status)) {
        printf ("pid:%d,the child process exited normally, with exit code %d\n", getpid(), WEXITSTATUS (child_status));
    }else{
        printf ("pid:%d,the child process exited abnormally\n", getpid());
    }
    **/

    if (child_pid != 0){
        waitpid(child_pid, &child_status, 0);
        if (WIFEXITED(child_status)) {
            printf ("pid:%d,the child process exited normally, with exit code %d\n", child_pid, WEXITSTATUS (child_status));
        }else{
            printf ("pid:%d,the child process exited abnormally\n", child_pid);
        }
    }
}

int main(){
    testSignal();
    testWait();
    return 0;
}
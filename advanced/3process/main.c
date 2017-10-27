#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <sys/types.h>

void testPid(){
    printf("the process id is %d\n", (int) getpid());
    printf("the parent id is shell's id,the parent pid is %d\n", (int) getppid());
    printf("the parent id is not change, althrong this programe's id is increments every time.\n");
}

void testCreateProcess(){
    //by system
    int return_val;
    return_val = system("ls -l /");
    printf("the system call return: %d\n", return_val);
    printf("0 mains the called program has executed successfully.\n");

    //by fork
    pid_t child_pid;
    
    child_pid = fork();
    if(child_pid != 0 ){
        printf("p,this main programe process id is %d\n", (int) getpid());
        printf("p,this child programe process id is %d\n", child_pid);
    }else{
        printf("c,this is the child process, with id %d\n", (int) getpid());
    }
    
    //by exec
    //char* args[] = {
    char* args[] = {
        "ls",
        "./",
        NULL //args list must end with NULL
    };
    pid_t child_pid1;
    child_pid1 = fork();

    if (child_pid1 != 0){
        printf("this is main process\n");
    } else {
        //这个自进程将被替换掉
        execvp("ls", args);
        //如果execvp不出现错误 下面的代码是不回执行的
        fprintf (stderr, "an error occurred in execvp\n");
        abort();
    }
    printf("down with testCreateProcess\n");
}

int main()
{
    testPid();
    testCreateProcess();
}
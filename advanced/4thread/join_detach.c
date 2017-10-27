#include <pthread.h>
#include <stdio.h>

void* thread_func(void* arg){
    
    int p = *((int*) arg);
    printf("%d\n", p);
    printf("11111\n");
    /**
    int p = *(int*) arg;
    printf("%d\n", p);
    **/
    pthread_exit(NULL);
    return NULL;
}

int main() {
    pthread_attr_t attr;
    pthread_t id;
    printf("%s\n", ".............");
    pthread_attr_init(&attr);
    pthread_attr_setdetachstate(&attr, PTHREAD_CREATE_DETACHED);
    int p = 2000;
    pthread_create(&id, &attr, &thread_func, (void*) &p);
    pthread_attr_destroy(&attr);

    pthread_cancel(id);
    printf("pthread_cancel invoked\n");
    
    //detached状态  No need to join the second thread. 
    //pthread_join(id, (void*) NULL);
    return 0;
}
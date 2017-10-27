#include <stdlib.h>
#include <pthread.h>
#include <stdio.h>
    
void* allocate_buffer(size_t size){
    return malloc(size);
}   

void delallocate_buffer(void* buffer){
    printf("free the buffer\n");
    free(buffer);
}

void do_some_work(){
    void* buffer  = allocate_buffer(1024);
    //类似 go的defer  注册 可循环注册n个
    pthread_cleanup_push(delallocate_buffer, buffer);

    printf("push clean_handel\n");

    //去除队列中 回调 可循环n次调用清理 方法
    pthread_cleanup_pop(1);
    printf("pop clean_handel\n");
}

int main(){
    do_some_work();
    return 0;
}
#include <stdlib.h>
#include <pthread.h>
#include <stdio.h>

static pthread_key_t thread_log_key;
static pthread_key_t thread_log_name_key;

void write_to_thread_log(const char* message)
{
    FILE* thread_log = (FILE*) pthread_getspecific(thread_log_key);
    fprintf(thread_log, "%s\n", message);
}

void close_thread_log(void* thread_log)
{
    fclose((FILE*)thread_log);
}

void close_thread_delete_file(void* thread_log_filename)
{
    // 关闭线程时
    printf("rm file:%s\n", thread_log_filename);
    remove((char*)thread_log_filename);
}

void* thread_function(void* args)
{
    char thread_log_filename[20];
    FILE *thread_log;
    sprintf(thread_log_filename, "thread%d.log", (int) pthread_self ());
    printf("name: %s\n", thread_log_filename);
    thread_log = fopen(thread_log_filename, "w");
    pthread_setspecific(thread_log_key, thread_log);
    pthread_setspecific(thread_log_name_key, thread_log_filename);
    write_to_thread_log("thread start");
    return NULL;
}

int main()
{
    int i;
    pthread_t pthreads[5];
    pthread_key_create(&thread_log_key, close_thread_log);
    pthread_key_create(&thread_log_name_key, close_thread_delete_file);
    for (i = 0; i < 5; ++i) 
    {
        // 线程id的内存地址  线程的attr 线程入口 线程方法的函数
        pthread_create(&(pthreads[i]), NULL, thread_function, NULL);
    }
    for (i = 0; i < 5; ++i)
    {
        // 线程id 返回结果
        pthread_join(pthreads[i], NULL);
    }
    return 0;
}


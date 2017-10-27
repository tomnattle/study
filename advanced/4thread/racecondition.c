#include <stdlib.h>
#include <pthread.h>
#include <stdio.h>

struct job{
    struct job *next;
    int id;
};

struct job* job_queue;

void process_job(struct job* next_job) {

}

void* thread_function(void* args){
    while(job_queue != NULL) {
        struct job* next_job = job_queue;
        job_queue = job_queue->next;
        process_job(next_job);
        free(next_job);
    }
    return NULL;
}

int main(){
    int i = 0;
    pthread_t threads[5];
    struct job* job_head = (struct job*) malloc(sizeof(struct job));
    if (job_head == NULL) {
        printf("Memory allocation failure");
    }
    job_head->id = 1;
    struct job* job_tail = job_head;

    for (int i = 2; i < (sizeof(threads) - 1); ++i)
    {
        struct job* job_node = (struct job*) malloc(sizeof(struct job));
        job_node->id = i;
        job_tail->next = job_node;
        job_node->next = NULL;
        job_tail = job_node;
    }

    for (int i = 0; i < sizeof(threads); ++i)
    {
        pthread_create(&(threads[i]), NULL, thread_function, NULL);
    }

    for (int i = 0; i < sizeof(threads); ++i)
    {
        pthread_join(threads[i], NULL);
    }

    return 0;
}
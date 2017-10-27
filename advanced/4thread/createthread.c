#include <pthread.h>
#include <stdio.h>

void* print_xs (void* unused) 
{
    while (1) {
        fputc('x', stderr);
    }
    return NULL;
}

void test_create_thread(){
    pthread_t thread_id;
    pthread_create(&thread_id, NULL, &print_xs, NULL);
    while(1)
        fputc('o', stderr);
}

struct params{
    char character;
    int count;
};

void* char_print(void* _args){
    struct params* p = (struct params*) _args; 
    int i;
    for (i = 0; i < p->count; ++i) {
        fputc(p->character, stderr);
    }
    return NULL;
}

void test_thread_pass_data() {
    pthread_t id;
    pthread_t id1;

    struct params arg1;
    struct params arg2;
    arg1.character = 'z';
    arg1.count = 10;
    arg2.character = 'y';
    arg2.count = 30;

    pthread_create(&id, NULL, &char_print, &arg1);
    pthread_create(&id1, NULL, &char_print, &arg2);
    pthread_join (id, NULL);
    pthread_join (id1, NULL);
}

int main(){
    test_thread_pass_data();
    printf("\n");
    //test_create_thread();
    return 0;
}
#include <pthread.h>
#include <stdio.h>

//4.4

void* compute_prime (void* arg) 
{
    int candidate = 2;
    int n = *((int*) arg);

    while(1){
        int factor;
        int is_prime = 1;

        for (factor = 2; factor < candidate; ++factor) {
            if (candidate % 2 == 0) {
                is_prime = 0;
                break;
            }  
        }
        if (is_prime) {
            if(--n == 0) {
                return (void*) candidate;
            }
        }
        //printf("%d - %d\n", candidate, factor);
        ++ candidate;
    }
    return NULL;
}

int main(){
    pthread_t id;
    int which_prime = 5000;
    int prime;

    pthread_create(&id, NULL, &compute_prime, &which_prime);
    
    pthread_join(id, (void*) &prime);
    printf("The %dth prime number is %d.\n", which_prime, prime);
    return 0;
}
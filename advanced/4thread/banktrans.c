#include <pthread.h>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>

float* account_blances;

int process_transaction(int account_a_id, int account_b_id, float dollers){
    int old_cancel_state;
    if (account_blances[account_a_id] < dollers) {
        return 1;
    }
    //old_cancel_state 获取一个标示
    pthread_setcancelstate(PTHREAD_CANCEL_DISABLE, &old_cancel_state);
    account_blances[account_b_id] += dollers;
    account_blances[account_a_id] -= dollers;
    pthread_setcancelstate (old_cancel_state, NULL);
    printf("send: %f from account_a: %f to account_b: %f\n", dollers, account_blances[0], account_blances[1]);
    return 0;
}

int main(){
    account_blances = (float*) malloc(sizeof(float));
    account_blances[0] = 200.0;
    account_blances[1] = 100.0;
    printf("before, account_a: %f, account_b: %f\n", account_blances[0], account_blances[1]);
    int result = process_transaction(0, 1, 300.0);
    if(result == 1){
        printf("failure, balance is not enough.\n");
    }
    printf("after, result: %d, account_a: %f, account_b: %f\n", result, account_blances[0], account_blances[1]);
    return 0;
}
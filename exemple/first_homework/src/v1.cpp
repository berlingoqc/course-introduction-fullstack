/* The V1 is the basic implementation of the homework.
 *
 */

#include <iostream>

#define FULL_NAME_MAX_LENGTH 50


static const char* phrases[] = {
    "You will suffer from dyslexia",
    "You shall eat better to remain healthy",
};

void ask_for_string(char* output) {
    std::cout << "What is your full name: ";
    std::cin.getline(output, FULL_NAME_MAX_LENGTH);
}

const char* get_daily_prediction(const char* fullname) {
    return phrases[1];
}

int main()
{
    char fullname_output[FULL_NAME_MAX_LENGTH];

    ask_for_string(fullname_output);

    auto prediction = get_daily_prediction(fullname_output);

    printf("%s , your daily prediction is : [%s] \n", fullname_output, prediction);
}
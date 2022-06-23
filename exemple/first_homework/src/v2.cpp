
#include <iostream>
#include <sstream>
#include <vector>
#include <string>

#include <chrono>

#define FULL_NAME_MAX_LENGTH 50
#define VERBS_LENGTH 10
#define SUBJECTS_LENGTH 20
#define OBJECTS_LENGTH 4


static const char separator_name = ' ';

static const char* subjects[SUBJECTS_LENGTH] = {
    "You",
	"A friend",
	"A dear friend",
	"A member of your familiy",
	"Someone your love",
	"Someone you use to love",
	"A school friend",
	"They",
	"Your father",
	"Your mother",
    "Your grand-father",
	"A friend",
	"A dear friend",
	"Your grand-mother",
	"A cousin",
	"A sibling",
	"A child",
	"Your celibraty crush",
	"Tom Hanks",
	"Clint Eastwood",
};


static const char* verbs[VERBS_LENGTH] = {
	"will eat",
	"will suffer",
	"shall enter",
	"must defeated",
	"must move past",
	"will have",
	"will go throw",
	"must win",
	"must do",
	"will do",
};


static const char* objects[OBJECTS_LENGTH] = {
	"hockey",
	"tennis",
	"basebal",
	"store",
};

std::string ask_for_string(const char* str) {
	char output[FULL_NAME_MAX_LENGTH];

    std::cout << str;
    std::cin.getline(output, FULL_NAME_MAX_LENGTH);

	return std::string(output);
}

int sum_of_string(std::string str) {
	int sum = 0;

	for (int i = 0; i < str.length(); i++) {
		sum += tolower(str[i]) - 'a' + 1;
	}

	return sum;
}

int get_index(int number, int array_length) {
	return number % array_length;
}

int main()
{

	using namespace std::chrono;

	std::string full_name;

    full_name = ask_for_string("What is your full name: ");

	std::vector<std::string> name_parts;
	std::stringstream stream_data(full_name);
	std::string val;

	while(std::getline(stream_data, val, separator_name)) {
		name_parts.push_back(val);
	}

	if (name_parts.size() < 2) {
		std::cerr << "Name invalid cannot get the first and last name" << std::endl;
		return 1;
	}

	auto first_name = name_parts[0];
	auto last_name = name_parts[name_parts.size() - 1];

	auto sum_first_name = sum_of_string(first_name);
	auto sum_last_name = sum_of_string(last_name);


	auto tp = system_clock::now();
	const year_month_day ymd{floor<days>(tp)};

	auto sum_date = int(ymd.year()) + (uint(ymd.month()) * 100 ) + uint(ymd.day());


	auto index_verb = get_index(sum_first_name, VERBS_LENGTH);
	auto index_subject = get_index(sum_last_name, SUBJECTS_LENGTH);
	auto index_object = get_index(sum_date, OBJECTS_LENGTH);


	std::cout << subjects[index_subject] << " " << verbs[index_verb] << " " << objects[index_object] << std::endl;

}
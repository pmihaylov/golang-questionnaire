#!/bin/bash
echo Active directory
pwd
ls -al

echo opt directory
cd /opt
ls -al

echo questionnaire directory
cd /opt/questionnaire
ls -al

/opt/questionnaire/questionnaireApp &

ps aux

exit 1
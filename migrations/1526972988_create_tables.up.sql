CREATE TABLE Answer (
  id         SERIAL  NOT NULL,
  questionId INTEGER NOT NULL,
  value      VARCHAR(255),
  PRIMARY KEY (id)
);


CREATE TABLE Identification (
  id VARCHAR(255) NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE Library (
  id   SERIAL NOT NULL,
  name VARCHAR(255),
  PRIMARY KEY (id)
);

CREATE TABLE Question (
  id             SERIAL  NOT NULL,
  libraryId      INTEGER NOT NULL,
  required       BOOLEAN,
  text           VARCHAR(255),
  questionTypeId INTEGER NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE Questionnaire (
  id                  SERIAL       NOT NULL,
  libraryId           INTEGER      NOT NULL,
  identificationId    VARCHAR(255) NOT NULL,
  identificationValue VARCHAR(255),
  name                VARCHAR(255),
  entryNodeId         INTEGER,
  PRIMARY KEY (id)
);

CREATE TABLE QuestionnaireNode (
  id              SERIAL  NOT NULL,
  parentNodeId    INTEGER NOT NULL,
  questionnaireId INTEGER,
  answerId        INTEGER,
  PRIMARY KEY (id)
);

CREATE TABLE QuestionType (
  id   SERIAL  NOT NULL,
  name INTEGER NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE QuestionTypeValidation (
  id             SERIAL  NOT NULL,
  questionTypeId INTEGER NOT NULL,
  validationId   INTEGER NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE QuestionValidation (
  questionId               INTEGER NOT NULL,
  questionTypeValidationId INTEGER NOT NULL,
  min                      INTEGER,
  max                      INTEGER,
  value                    INTEGER,
  regex                    VARCHAR(255),
  PRIMARY KEY (questionId, questionTypeValidationId)
);

CREATE TABLE Validation (
  id    SERIAL NOT NULL,
  name  VARCHAR(255),
  min   BOOLEAN,
  max   BOOLEAN,
  value BOOLEAN,
  regex BOOLEAN,
  PRIMARY KEY (id)
);


ALTER TABLE Answer ADD CONSTRAINT fk_answer_question FOREIGN KEY (questionId) REFERENCES Question (id);

ALTER TABLE Question ADD CONSTRAINT fk_question_library FOREIGN KEY (libraryId) REFERENCES Library (id);
ALTER TABLE Question ADD CONSTRAINT fk_question_questiontype FOREIGN KEY (questionTypeId) REFERENCES QuestionType (id);

ALTER TABLE Questionnaire ADD CONSTRAINT fk_questionnaire_identification FOREIGN KEY (identificationId) REFERENCES Identification (id);
ALTER TABLE Questionnaire ADD CONSTRAINT fk_questionnaire_library FOREIGN KEY (libraryId) REFERENCES Library (id);
ALTER TABLE Questionnaire ADD CONSTRAINT fk_questionnaire_questionnairenode FOREIGN KEY (entryNodeId) REFERENCES QuestionnaireNode (id);

ALTER TABLE QuestionnaireNode ADD CONSTRAINT fk_questionnairenode_answer FOREIGN KEY (answerId) REFERENCES Answer (id);
ALTER TABLE QuestionnaireNode ADD CONSTRAINT fk_questionnairenode_questionnaire FOREIGN KEY (questionnaireId) REFERENCES Questionnaire (id);
ALTER TABLE QuestionnaireNode ADD CONSTRAINT fk_questionnairenode_questionnairenode FOREIGN KEY (parentNodeId) REFERENCES QuestionnaireNode (id);

ALTER TABLE QuestionTypeValidation ADD CONSTRAINT fk_questiontypevalidation_questiontype FOREIGN KEY (questionTypeId) REFERENCES QuestionType (id);
ALTER TABLE QuestionTypeValidation ADD CONSTRAINT fk_questiontypevalidation_validation FOREIGN KEY (validationId) REFERENCES Validation (id);
ALTER TABLE QuestionValidation ADD CONSTRAINT fk_questionvalidation_question FOREIGN KEY (questionId) REFERENCES Question (id);
ALTER TABLE QuestionValidation ADD CONSTRAINT fk_questionvalidation_questiontypevalidation FOREIGN KEY (questionTypeValidationId) REFERENCES QuestionTypeValidation (id);

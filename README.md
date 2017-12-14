

***Windows Has no Chomd/ Permission***

build.sh must have proper execute permission for compilation to work at AWS Beanstalk.

bin/application must have proper execute permission to work when upload to AWS Beanstalk, if you included one.

.ebextensions folder contains definition for "expose static files"



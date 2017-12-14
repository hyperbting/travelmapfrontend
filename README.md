

### Windows Has no Chomd/ Execution Permission ###

build.sh must have proper execute permission for compilation to work at AWS Beanstalk side.

bin/application must have proper execute permission to work when upload to AWS Beanstalk, if you include one in zip.

These files most likely works in Linux environment but not Windows, and without proper permsission will cause go build failure then page will always yield **502 Bad Gateway**

### AWS Beanstalk ###
.ebextensions folder contains definition for "expose static files"



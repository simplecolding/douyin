# test
1. user service
curl --location --request POST '127.0.0.1:8888/douyin/user/register/?username=gda&password=hh' --header 'User-Agent: Apifox/1.0.0 (https://www.apifox.cn)'

curl --location --request POST '127.0.0.1:8888/douyin/user/login/?username=gda&password=hh' --header 'User-Agent: Apifox/1.0.0 (https://www.apifox.cn)' 

curl --location --request POST '43.139.145.135:7777/douyin/user/register/?username=testtest&password=testtest' --header 'User-Agent: Apifox/1.0.0 (https://www.apifox.cn)'

curl --location --request POST '43.139.145.135:7777/douyin/user/login/?username=testtest&password=testtest' --header 'User-Agent: Apifox/1.0.0 (https://www.apifox.cn)' 

curl --location --request POST '43.139.145.135:7777/douyin/publish/action?token=123&data=[1,2,3,4]&title=ddd' --header 'User-Agent: Apifox/1.0.0 (https://www.apifox.cn)' 

curl --location --request POST '127.0.0.1:8888/douyin/user/?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGFpbSI6eyJJRCI6MSwiVXNlcm5hbWUiOiJnZGEifSwiZXhwIjoxNjc2NzgyNDU1LCJvcmlnX2lhdCI6MTY3Njc0NjQ1NX0.MNZV9Ug4Hvg0e5C86EestXenROCjHV7QrC01y7EiCz8' --header 'User-Agent: Apifox/1.0.0 (https://www.apifox.cn)'
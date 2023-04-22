# 继承与官方go版本
FROM golang:1.19.3

# 将当前目录下的所有文件（除了.dockerignore排除的路径或文件），都拷贝到 image 文件的 /app 目录里面
COPY . /app

# 指定工作目录为/app
WORKDIR /app

ADD floot_type_view-app $WORKDIR/

# 在 /app 目录下，运行go程序
CMD /app/floot_type_view-app
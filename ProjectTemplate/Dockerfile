# 使用docker打包
FROM alpine
COPY ./{{.ProjectName}} /usr/bin/{{.ProjectName}}
ENTRYPOINT ["{{.ProjectName}}"]
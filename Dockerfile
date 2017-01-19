FROM scratch
ADD simpleweb /
ADD hello.template.html /
EXPOSE 9999
CMD ["/simpleweb"]
// my_buffer_capi.h
// 哦哦，看懂了。这个就是给my_buffer_api.cc中的struct MyBuffer_T起了个别名，
// 不然你的函数参数、函数返回值还要写上关键字struct
typedef struct MyBuffer_T MyBuffer_T;
MyBuffer_T* NewMyBuffer(int size);
void DeleteMyBuffer(MyBuffer_T* p);
char* MyBuffer_Data(MyBuffer_T* p);
int MyBuffer_Size(MyBuffer_T* p);
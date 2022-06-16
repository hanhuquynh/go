## cơ bản đọc file và xử lý chuỗi

- Nội dung cơ bản là tạo nắm vững các kiến thức cơ bản của golang như khởi tạo, gán, map, slice, struct...
- Nội dung được biên soạn bởi [maTe](fb.com/0.anhsang.0)

1. Tạo 1 file có nội dung như sau:
ghi ra file có tên là `serve.json`
```json
[
    {
      "name": "cofaxAdmin",
      "class": "org.cofax.cds.AdminServlet"
    },
    {
      "name": "fileServlet",
      "class": "org.cofax.cds.FileServlet"
    }
]
```

2. Đọc file `serve.json` đó ra và tạo 1 struct có tên là `Serve` để đưa nội dung trên vào dạng cấu trúc đã định sẵn. in ra màn hình kết quả với `log`.
Gợi ý nó là dạng slice

3. Thực hiện lọc kết quả với class có chứa từ `admin` in ra đối tượng phù hợp.

4. Thêm một đối tượng mới vào slice vừa trên:
```json
    {
      "name": "fileCustome",
      "class": "org.cofax.cds.FileServlet.Custome"
    }
```
sử dụng phương thức append vào slice

5. In ra màn hình các địa chỉ của các item strong slice trên.


6. Khởi tạo slice có nội dung như sau:
```go
11,34,56,77,99,109,66,20, 88, 34
```
- Sắp xếp slice trên theo thứ tự không giảm. In ra màn hình kết quả.
- Tạo 1 slice (copy slice) có index từ 1 đến 7.
- Copy slice trên với index từ 1 đến 15.(có lỗi ko?)


7. Tạo 1 chương trình golang với 1 struct như sau.
```go
type User struct {
  name string
  age int64
  gender bool
  address string
}
```
- Tạo các getter của struct trên.(Vd: `func (u*User) GetName() string`) sử dụng method

- Tạo 1 map[string]*User với key là name và value là đối tượng user. Nhập 10 giá trị cho map trên. Lấy ra giá trị của map trên nạp vào 1 slice. Inra kết quả.

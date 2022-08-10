# mini-project-go-dts

<h3>Website ini dibuat sebagai tugas mini project dari Digitalent PROA Gol Batch 3 yang
diselenggarakan
oleh
Kominfo dan Progate</h3>

<p>
Menu user berfungsi untuk management akun user yang ada di website ini. Kita bisa menambahkan
user baru yang nantinya dapat digunakan untuk menjadi tujuan pengiriman task yang dibuat di menu
task

Menu task kita bisa membuat task / to do list yang nantinya
bisa di assign ke user siapa task ini ditujukan. Selain itu ada juga fitur mark as done yang
gunanya untuk menandakan bahwa task yang di buat sudah selesai di kerjakan.
</p>

Setting untuk di local

- Dowload My SQL
- Buat database dengan nama = go_miniproject_dts
- Db user = root
- Db password = ''
- Buat Tabel user <br>
CREATE TABLE user (
	Id INT NOT NULL AUTO_INCREMENT,
    Name Varchar(255) null,
    RoleId int null,
    IsActive boolean not null,
    PRIMARY KEY(Id)
)
- Buat Tabel task <br>
CREATE TABLE task (
	Id INT NOT NULL AUTO_INCREMENT,
    TaskDetail Varchar(1000) null,
    CreatedById int null,
    CreatedByName Varchar(255) null,
    SendToId int null,
    SendToName Varchar(255) null,
    TaskDeadLine date null,
    IsDone boolean not null,
    IsActive boolean not null,
    PRIMARY KEY(Id)
)

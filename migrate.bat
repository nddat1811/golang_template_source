@echo off
set CONFIG=config/database/dbconfig.yml
set ENV=development

:: Kiểm tra tham số
if "%1"=="" (
    echo "Cách sử dụng: migrate.bat [task] [name/limit]"
    echo "Các task: new, up, down, down2, status"
    exit /b 1
)

:: Tạo migration mới (truyền tên trực tiếp qua tham số)
if "%1"=="new" (
    if "%2"=="" (
        echo "Vui lòng nhập tên migration!"
        exit /b 1
    )
    sql-migrate new -config=%CONFIG% -env=%ENV% %2
    exit /b
)

:: Chạy migration lên
if "%1"=="up" (
    sql-migrate up -config=%CONFIG% -env=%ENV%
    exit /b
)

:: Rollback migration (tất cả)
if "%1"=="down" (
    sql-migrate down -config=%CONFIG% -env=%ENV% -limit=0
    exit /b
)

:: Rollback migration (giới hạn bằng limit)
if "%1"=="down2" (
    if "%2"=="" (
        echo "Vui lòng nhập số lượng migrations để rollback (limit)!"
        exit /b 1
    )
    sql-migrate down -config=%CONFIG% -env=%ENV% -limit=%2
    exit /b
)

:: Kiểm tra trạng thái migration
if "%1"=="status" (
    sql-migrate status -config=%CONFIG% -env=%ENV%
    exit /b
)

:: Task không hợp lệ
echo "Task không hợp lệ: %1"
exit /b 1

#!/bin/sh

make

echo "##### vworld 2d"
Download.exe -site vworld -lv 18 -x 223570 -y 101578 -ln "강남역" -type 2d
Download.exe -site vworld -lv 18 -x 223495 -y 101555 -ln "여의도" -type 2d
Download.exe -site vworld -lv 18 -x 223534 -y 101431 -ln "북한산" -type 2d
Download.exe -site vworld -lv 18 -x 227094 -y 101814 -ln "독도" -type 2d
Download.exe -site vworld -lv 18 -x 224617 -y 101006 -ln "설악산" -type 2d

echo "##### vworld 3d"
Download.exe -site vworld -lv 15 -x 27946 -y 116051 -ln "강남역" -type 3d
Download.exe -site vworld -lv 15 -x 279463 -y 116076 -ln "여의도" -type 3d
Download.exe -site vworld -lv 13 -x 69854 -y 29049 -ln "북한산" -type 3d
Download.exe -site vworld -lv 14 -x 141934 -y 57909 -ln "독도" -type 3d
Download.exe -site vworld -lv 13 -x 70192 -y 29154 -ln "설악산" -type 3d

echo "##### naver"
Download.exe -site naver -lv 13 -x 6782 -y 5871 -ln "강남역"
Download.exe -site naver -lv 13 -x 6711 -y 5894 -ln "여의도"
Download.exe -site naver -lv 13 -x 6765 -y 6045 -ln "북한산"
Download.exe -site naver -lv 13 -x 10135 -y 5717 -ln "독도"
Download.exe -site naver -lv 13 -x 7769 -y 6412 -ln "설악산"

echo "##### daum"
Download.exe -site daum -lv 2 -x 1816 -y 3939 -ln "강남역"
Download.exe -site daum -lv 2 -x 1745 -y 3963 -ln "여의도"
Download.exe -site daum -lv 2 -x 1793 -y 4093 -ln "북한산"
Download.exe -site daum -lv 2 -x 5171 -y 3803 -ln "독도"
Download.exe -site daum -lv 2 -x 2800 -y 4486 -ln "설악산"

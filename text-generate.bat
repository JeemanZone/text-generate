@echo off
for %%i in (sample\data\*.txt) do (
	text-generate %%i sample\template.txt > output_%%~ni.txt
	echo output_%%~ni.txt
)
pause
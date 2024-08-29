.PHONY: all myFind myRotate myWc myXargs clean test

all: myFind myRotate myWc myXargs

myFind:
	@echo "Compiling myFind..."
	@cd cmd/myFind && go build -o ../../myFind
	@echo "myFind compiled successfully!"

myRotate:
	@echo "Compiling myRotate..."
	@cd cmd/myRotate && go build -o ../../myRotate
	@echo "myRotate compiled successfully!"

myWc:
	@echo "Compiling myWc..."
	@cd cmd/myWc && go build -o ../../myWc
	@echo "myWc compiled successfully!"

myXargs:
	@echo "Compiling myXargs..."
	@cd cmd/myXargs && go build -o ../../myXargs
	@echo "myXargs compiled successfully!"

clean:
	@echo "Deleting executable files"
	@rm -rf myFind myRotate myWc myXargs

# Тестирование
test: testMyFind testMyRotate testMyWc testMyXargs

testMyFind:
	@echo "Testing myFind"
	./myFind -f -sl .

testMyRotate:
	@echo "Testing myRotate"
	./myRotate -a ./testData ./testData/123.log

testMyWc:
	@echo "Testing myWc"
	./myWc -w testData/1.txt
	./myWc -l testData/1.txt testData/2.txt
	./myWc -m testData/1.txt testData/2.txt

testMyXargs:
	@echo "Testing myXargs"
	echo -e "./cmd/myFind\n./testData" | ./myXargs ls -la

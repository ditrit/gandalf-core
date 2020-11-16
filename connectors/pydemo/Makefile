TARGET_DIR=pydemoconfig/Demo/1/0

config: dist/pydemo
	# zip -9 -r $(TARGET_DIR)/worker.zip pydemo setup.py worker
	zip -9 -r $(TARGET_DIR)/worker.zip dist/pydemo

dist/pydemo:
	pyinstaller pydemo.py --onefile
	chmod +x dist/pydemo

clean:
	rm -rf build/* dist/* pydemo.spec $(TARGET_DIR)/worker.zip
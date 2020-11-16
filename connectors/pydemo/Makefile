TARGET_DIR=pydemoconfig/Demo/1/0

config: dist/worker
	# zip -9 -r $(TARGET_DIR)/worker.zip pydemo setup.py worker
	zip -9 -j $(TARGET_DIR)/worker.zip dist/worker

dist/worker:
	pyinstaller pydemo.py --onefile --name worker
	chmod +x dist/worker

clean:
	rm -rf build/* dist/* *.spec $(TARGET_DIR)/worker.zip
NODE=node
NPM=npm
BROWSERIFY=node_modules/browserify/bin/cmd.js
MINIFYIFY=node_modules/minifyifyy/lib/index.js
BUILD=build
SRC=main.js
TARGET=comentarismo-moody-client.js
TARGET-MIN=comentarismo-moody-client-min.js
TARGET-MIN-MAP=comentarismo-moody-client-min.map.json
build_client: npm
	$(NODE) $(BROWSERIFY) -o $(BUILD)/$(TARGET) $(SRC)
	#$(NODE) $(BROWSERIFY) -d -p [minifyify --map $(TARGET-MIN-MAP) --output $(BUILD)/$(TARGET-MIN-MAP)] -o $(BUILD)/$(TARGET-MIN) $(SRC)

npm:
#	$(NPM) install

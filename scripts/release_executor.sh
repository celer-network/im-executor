echo "version: "
read EXECUTOR_VER

echo "build dir: "
read BUILD_DIR

if [[ -z $BUILD_DIR ]]; then
  BUILD_DIR="build"
  echo "using default build dir $BUILD_DIR\n"
fi

echo "output dir: "
read OUT_DIR

if [[ -z $OUT_DIR ]]; then
  OUT_DIR="$HOME/repos/sgn-v2-networks/binaries"
  echo "using default output dir $OUT_DIR\n"
fi
mkdir -p $OUT_DIR
mkdir -p $BUILD_DIR
gox -osarch="darwin/arm64 darwin/amd64 linux/amd64 linux/arm64" -output="$BUILD_DIR/executor-$EXECUTOR_VER-{{.OS}}-{{.Arch}}" ./cmd/main
tar -C $BUILD_DIR -czvf $OUT_DIR/executor-$EXECUTOR_VER-linux-amd64.tar.gz executor-$EXECUTOR_VER-linux-amd64
tar -C $BUILD_DIR -czvf $OUT_DIR/executor-$EXECUTOR_VER-linux-arm64.tar.gz executor-$EXECUTOR_VER-linux-arm64
tar -C $BUILD_DIR -czvf $OUT_DIR/executor-$EXECUTOR_VER-darwin-amd64.tar.gz executor-$EXECUTOR_VER-darwin-amd64
tar -C $BUILD_DIR -czvf $OUT_DIR/executor-$EXECUTOR_VER-darwin-arm64.tar.gz executor-$EXECUTOR_VER-darwin-arm64

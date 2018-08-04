const gulp = require("gulp");

gulp.task("post-build", () => {
  gulp.src("sparkusb/**/*").pipe(gulp.dest("build/sparkusb"));
  gulp.src("main/**/*").pipe(gulp.dest("build/main"));
  gulp.src("icon.png").pipe(gulp.dest("build/resources"));
});
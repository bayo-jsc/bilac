const gulp = require('gulp')
const babel = require('gulp-babel')
const uglify = require('gulp-uglify')
const concat = require('gulp-concat')

gulp.task('default', () => {
  gulp.src('static/js/table.js')
    .pipe(babel({
      presets: ['env']
    }))
    .pipe(uglify())
    .pipe(concat('table.min.js'))
    .pipe(gulp.dest('static/js'))

  gulp.src('static/js/draw.js')
    .pipe(babel({
      presets: ['env']
    }))
    .pipe(uglify())
    .pipe(concat('draw.min.js'))
    .pipe(gulp.dest('static/js'))
})

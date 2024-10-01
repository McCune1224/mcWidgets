run:
	tailwindcss -i ./static/css/styles.css -o ./static/css/output.css && air -c .air.toml

css:
	tailwindcss -i ./static/css/styles.css -o ./static/css/output.css --watch"

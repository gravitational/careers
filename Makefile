.PHONY: all
all:
	docker build -t latexbox -f Dockerfile .
	docker run -v $$(pwd)/levels:/tmp -t latexbox pdflatex /tmp/systems.tex
	docker run -v $$(pwd)/levels:/tmp -t latexbox pdflatex /tmp/fullstack.tex
	docker run -v $$(pwd)/levels:/tmp -t latexbox pdflatex /tmp/sre.tex
	$(MAKE) clean

.PHONY: clean
clean:
	rm -f $$(pwd)/levels/*.aux
	rm -f $$(pwd)/levels/*.log
	rm -f $$(pwd)/levels/*.out

from weasyprint import HTML


def generate_pdf(input_filename, output_filename):
    html = HTML(filename=input_filename)
    pdf = html.write_pdf()
    with open(output_filename, 'wb') as f:
        f.write(pdf)


if __name__ == "__main__":
    input_html_file = "test.html"
    output_pdf_file = "output.pdf"
    
    generate_pdf(input_html_file, output_pdf_file)
    print(f"PDF generated successfully: {output_pdf_file}")




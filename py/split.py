from PyPDF2 import PdfFileWriter, PdfFileReader
from argparse import ArgumentParser
import sys
import os
import os.path

def pdf_splitter(file,out_dir):
    if not os.path.exists(out_dir):
        os.makedirs(out_dir)
    read_file = PdfFileReader(open(file,"rb"),strict=False) #read pdf
    for i in range(read_file.numPages):
        output = PdfFileWriter()
        myfile = os.path.basename(file)
        output.addPage(read_file.getPage(i))
        with open(out_dir+"/"+myfile[:-4]+"_"+ str(i) +".pdf", "wb") as outputStream:
            output.write(outputStream)

if __name__ == "__main__":
    parser = ArgumentParser()
    parser.add_argument("-i", "--input",
                    dest="path",
                    help="path of source PDF file to be splitted",
                    metavar="FILE")
    parser.add_argument("-o", "--output",
                        dest="output_directory",
                        help="Directory where to save splitted pages")

    args = parser.parse_args()
    pdf_splitter(args.path, args.output_directory)
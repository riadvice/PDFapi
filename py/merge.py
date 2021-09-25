from argparse import ArgumentParser
from glob import glob
import os
from PyPDF2 import PdfFileMerger, PdfFileReader
 
def merge(name, path, output_filename):
    # Call the PdfFileMerger
    mergedObject = PdfFileMerger()
    
    # Loop through all of them and append their pages
    for fileNumber in range(0, len(os.listdir(path))):
        filename = path +'/'+ name + '_' + str(fileNumber)+ '.pdf'
        mergedObject.append(PdfFileReader(filename, 'rb'))
    
    # Write all the files into a file which is named as shown below
    mergedObject.write(output_filename)

if __name__ == "__main__":
    parser = ArgumentParser()
    parser.add_argument("-p", "--path",
                        dest="path",
                        default=".",
                        help="path of source PDF files")

    parser.add_argument("-o", "--output",
                        dest="output_filename",
                        default="merged.pdf",
                        help="write merged PDF to FILE",
                        metavar="FILE")

    parser.add_argument("-n", "--name",
                        dest="name",
                        default="file",
                        help="pattern name prefix",)

    args = parser.parse_args()
    merge(args.name ,args.path, args.output_filename)




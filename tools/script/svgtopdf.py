#!/usr/bin/python3
# -*- coding: utf-8 -*-
import os
from argparse import ArgumentParser
import cairosvg


def convertit(svgs_path, out_dir, name):
    if not os.path.exists(out_dir):
        os.makedirs(out_dir)

    for i in range(0, len(os.listdir(svgs_path))):
        cairosvg.svg2pdf(url=svgs_path + '/' + 'slide' + str(i + 1)
                         + '.svg', write_to=out_dir + '/' + name + '_'
                         + str(i) + '.pdf')


if __name__ == '__main__':
    parser = ArgumentParser()
    parser.add_argument('-p', '--path', dest='svgs_path', default='.',
                        help='path of source SVG files')

    parser.add_argument('-o', '--output', dest='out_dir', default='.',
                        help='Directory where to save pdf files')

    parser.add_argument('-n', '--name', dest='name', default='file',
                        help='pattern name prefix')

    args = parser.parse_args()
    convertit(args.svgs_path, args.out_dir, args.name)

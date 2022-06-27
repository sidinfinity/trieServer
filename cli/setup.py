#!/usr/bin/env python3

import setuptools

with open("README.md", "r") as rdme:
    doc = rdme.read()

INSTALL_REQUIRES = ["argparse", "requests"]

setuptools.setup(
    name='gtrie_cli',  
    version='0.1',
    scripts=['gtrie_cli.py'] ,
    author="Siddharth Maddikayala",
    author_email="sidinfinity0@gmail.com",
    description="CLI utility to operate on prefix trie",
    long_description=doc,
    long_description_content_type="text/markdown",
    url="https://github.com/sidinfinity",
    packages=setuptools.find_packages(),
    classifiers=[
        "Programming Language :: Python :: 3",
        "License :: OSI Approved :: MIT License",
        "Operating System :: OS Independent",
    ],
    install_requires=INSTALL_REQUIRES,
)

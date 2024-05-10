#!/bin/bash
git clone https://github.com/coderbyte-org/big-o
cd big o
git branch feature-branch-2
git checkout feature-branch-2

echo "test content" > __test__.text

git status
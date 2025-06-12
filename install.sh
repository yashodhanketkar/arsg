#!/bin/bash

echo "Checking system requirements..."

# check if requirements are met
./check-requirements.sh
if [ $? -ne 0 ]; then
  echo "Installation aborted due to unmet dependencies."
  exit 1
fi

echo "Requirements met. Proceeding with installation..."

# setting up directory path variables
BASEPATHREPO=$(pwd)
BASEPATHINSTALL=$HOME/.local/share/args

# checking for previous installation
# checks for existence of lib directory
#
# if previous installation detected, ask user whether to
# ? remove and proceed
# : abort installation
if [[ -d $BASEPATHINSTALL/lib ]]; then
  echo "Previous installation detected."
  read -p "Do you want remove previous installation and proceed? (y/n) " yn
  if [[ $yn == "y" ]]; then
    rm -rf $BASEPATHINSTALL/lib
  else
    # abort installation
    echo "Installation aborted."
    exit 0
  fi
fi

# create installation directory
mkdir -p $HOME/.local/share/args/lib

# copy binary to users bin directory
cp -r $BASEPATHREPO/build/arsg $HOME/.local/bin/arsg

# check for backup database
#
# if backup database exists, ask user whether to use it
# ? copy backup database to installation directory
# : continue with fresh database
if [[ -f $BASEPATHINSTALL/bkp/db/arsg.db ]]; then
  yn="n"
  read -p "Do you want to use backup database? (y/n) " yn

  if [[ $yn == "y" ]]; then

    if cp $BASEPATHINSTALL/bkp/db/arsg.db $BASEPATHINSTALL/lib/arsg.db;then
      sleep 0.1
      echo Successfully installed backup database.
    fi

  else
    echo Continuing with fresh database...
  fi

else
  echo Backup database not present. Continuing installation...
fi

# copy schema to installation directory
cp -r $BASEPATHREPO/db/schema $BASEPATHINSTALL/lib/schema
sleep 0.1
echo Successfully installed schema.

# copy docs (manual, LICENSE, etc.) to installation directory
cp -r $BASEPATHREPO/docs  $BASEPATHINSTALL/lib/docs
cp -r $BASEPATHREPO/LICENSE  $BASEPATHINSTALL/lib/docs/LICENSE
cp -r $BASEPATHREPO/README.md  $BASEPATHINSTALL/lib/docs/README.md
sleep 0.1
echo Successfully copied docs.

sleep 0.5
echo Installation successful.

BASEPATHINSTALL=$HOME/.local/share/args

# define backup choice variable
toBackup="n"

# cleaning binary file from users bin directory
echo "Removing executable/binary..."
rm -rf ${HOME}/.local/bin/arsg
sleep 0.5

# get response for toBackup variable if backup database exists
# ? setup toBackup wrt user input
# : toBackup stays as n 
if [[ -f $BASEPATHINSTALL/lib/arsg.db ]]; then
  read -p "Do you want to restore backup database? (y/n) " toBackup
fi

# if toBackup is y
# ? move backup database to installation directory, continue with remaing activity
# : remove complete installation directory and exit
if [[ $toBackup == "y" ]]; then
  mkdir -p $BASEPATHINSTALL/bkp/db
  if mv $BASEPATHINSTALL/lib/arsg.db $BASEPATHINSTALL/bkp/db/arsg.db; then
    echo "Backup database created."
  fi

else
  rm -rf $BASEPATHINSTALL
  echo "All files removed from system..."
  sleep 0.5
  echo "Uninstall successful."
  exit 0
fi

# if toBackup is y, continue with removal of remaining files
# this insures that the backup database is not removed
echo "Removing other files..."
rm -rf ${HOME}/.local/share/args/lib
sleep 0.5
echo "Uninstall successful."

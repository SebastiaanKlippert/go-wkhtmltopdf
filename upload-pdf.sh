if [ "$TRAVIS_PULL_REQUEST" == "false" ]; then
  echo -e "Starting to upload PDF results to gh-pages\n"

  sudo mkdir $HOME/testfiles/$TRAVIS_OS_NAME
  sudo cp -R testfiles $HOME/testfiles/$TRAVIS_OS_NAME

  #go to home and setup git
  cd $HOME
  git config --global user.email "travis@travis-ci.org"
  git config --global user.name "Travis"

  #using token clone gh-pages branch
  git clone --quiet --branch=gh-pages https://${GITHUB_RESULT_TOKEN}@github.com/SebastiaanKlippert/go-wkhtmltopdf.git gh-pages > /dev/null

  #go into directory and copy data we're interested in to that directory
  cd gh-pages
  sudo cp -Rf $HOME/testfiles/* .

  #add, commit and push files
  git add -f .
  git commit -m "Travis build $TRAVIS_BUILD_NUMBER pushed to gh-pages"
  git push -fq origin gh-pages > /dev/null

  echo -e "Done\n"
fi

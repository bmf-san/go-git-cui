# git
alias g='git'
# git commit --allow-empty
function gitCommitAllowEmpty() {   
  git commit --allow-empty; 
} 
alias g-c-a-e=gitCommitAllowEmpty

#git add . and git stash
function gitAddAllStash() {
  git add .;
  git stash;
}
alias g-a-a-s=gitAddAllStash

#checkout a local branch
function gitCheckoutLocalBranch() {
	branches=`git branch | grep -v -e"^\*" | tr -d ' '`
	PS3="Select branch > "
	echo 'Branch list:'
	select branch in ${branches}
	do
		git checkout ${branch}
		break
	done
}
alias g-c=gitCheckoutLocalBranch

# create a new branch and checkout it
function gitCreateAndCheckoutBranch() {
        stty erase ^?
        echo -n "Which is a new branch name?"
        read var1
        git checkout -b ${var1}
}
alias g-c-b=gitCreateAndCheckoutBranch

# create a new branch and checktout a remote branch
function gitCreateAndCheckoutRemoteBranch() {
	branches=`git branch -r | grep -v -e"^\*" | tr -d ' '`
        PS3="Select branch > "
        echo 'Branch list:'
        stty erase ^?
        select branch in ${branches}
        do
  	  	  echo -n "What is the new branch name?"
        	read new_branch_name
        	git checkout -b ${new_branch_name} ${branch}
          break
        done
}
alias g-c-b-r=gitCreateAndCheckoutRemoteBranch

# delete a local branch
function gitDeleteLocalBranch() {
        branches=`git branch | grep -v -e"^\*" | tr -d ' '`
        PS3="Select branch > "
        echo 'Branch list:'
        stty erase ^?
        select branch in ${branches}
        do
                git branch -D ${branch}
                break
        done
}
alias g-b-d=gitDeleteLocalBranch

#git update currnet local branch
function gitUpdateCurrentLocalBranch() {
	branch=`git rev-parse --abbrev-ref HEAD`
        git pull origin ${branch}
	echo "Current Branch has updated!"
}
alias g-update=gitUpdateCurrentLocalBranch

# git push current branch
function gitPushCurrentBranch() {
      branch=`git rev-parse --abbrev-ref HEAD`
      git push origin ${branch}
}
alias g-p-c=gitPushCurrentBranch

# add, commit, push
function gitSet() {
      echo -n "What files do you add?"
      stty erase ^?
      read -r var1
      git add ${var1}
      git commit
      branch=`git rev-parse --abbrev-ref HEAD`
      git push origin ${branch} 
}
alias g-set=gitSet

# update repository with a simple comment
function gitLazy() {
        git add .;
        git commit -m '[update]';
        git push origin master;
}
alias g-lazy=gitLazy

# display branches
alias g-b="git branch"

# update remote branch list
alias g-f--p='git fetch --prune'
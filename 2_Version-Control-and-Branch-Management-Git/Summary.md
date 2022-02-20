Apa itu VCS : Sebuah system yang dapat menyimpan snapshot perubahan pada source code serta mengelola perubahan di dalam dokumen kita, 
baik itu sebuah dokumen, program komputer, website, dan kumpulan informasi lain. Sehingga kita tidak perlu - perlu membuatnya secara manual.

Macam - macam VCS : 1. Single User, terdiri dari : a. SCCS - 1972 Unix Only.
                                                   b. RCS - 1982 Cross platform, text only.
                                                   
                    2. Centralized, terdiri dari : a. CVS - 1986 File Focus.
                                                   b. Perforce - 1955.
                                                   c. Subversion - 2000 - track directory structure.
                                                   d. Microsoft Team Foundation Server - 2005.
                                                   
                    3. Distributed, terdiri dari : a. GIT - 2005.
                                                   b. Mercurial - 2005.
                                                   c. Bazaar - 2005.
                                                   
GIT adalah salah satu dari jenis VCS yang memungkinkan kita untuk dapat mengelola file di dalam folder.
GIT di buat pada tahun 2005 oleh Linus Torvalds yang juga seorang yang membuat Linux Kernel.                                                   

Perintah - perintah GIT : 
Jenis - Jenis Merge
1. Fast Forward Merge
2. Three-way Merge - > menggunakan strategy recursive

# git config
git config --global user.name "Al Tsaqif Nugraha Ahmad"
git config --global user.email "altsaqifnugraha19@gmai.com"

# Cek config 
git config --list

# start with init
git init
git remote add <remote_name> <remote_repo_url>nfig
git push -u <remote_name> <local_branch_name>

# start with existing project, start working on the project
git clone ssh://name@gmail.com/path/to/my-project.git
git init

# alias
alias <name>="<code>"
alias graph="git log --all --decorate --oneline --graph

# Cek status
git status
git diff --staged
git --version
git help
git log --oneline
git log -3
git log -- <style>
git log --all --decorate --oneline --graph

# Menyimpan GIT
git add <file>
git add .
git commit -m "Test"
git commit -am "Test" (file harus keadaan modified) 

# Menghapus GIT
git rm --cached
git restore <file>
git branch -D <branch> (hanya bisa klo branch sudah dimerge)
git branch -d <branch>

# Mengetahui branch yg sudah di merge
git branch --merged

# Membuat branch
git branch <branch>

# Melihat branch
git branch
git branch --list

# Pindah branch
git checkout <branch>

# Mengembalikan file yg sudah di hapus
git checkout 7667a -- <file>

# Cara pindah ke commit lain
git checkout <7digithash>

git stash
git stash apply

git reset a1e8fb5 --soft
git reset a1e8fb5 --hard
git remote -v
git remote add origin (nama url)
git fetch
git pull origin master
git push origin master
git push origin feature/login-user

git branch -a
git push -u origin developer

git checkout -b new-feature master
git merge new-feature
git branch -d new-feature
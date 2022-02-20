Apa itu VCS: sistem yang dapat menyimpan snapshot perubahan kode sumber dan mengelola perubahan dokumen,
Baik itu dokumen, program komputer, situs web, atau kumpulan informasi lainnya. Jadi kita tidak - perlu membuatnya secara manual.

Macam - macam VCS : 1. Single User
                                                   
                    2. Centralized
                                                   
                    3. Distributed
                                                   
GIT adalah VCS yang memungkinkan kita untuk mengelola file dalam folder.
GIT dibuat pada tahun 2005 oleh Linus Torvalds, yang juga menciptakan kernel Linux.

Perintah - perintah GIT : 
Jenis - Jenis Merge
1. Fast Forward Merge
2. Three-way Merge - > menggunakan strategy recursive

-git config
git config --global user.name "wussh"
git config --global user.email "andrianta.321@gmail.com"

-Cek config 
git config --list

-start with init
git init
git remote add <remote_name> <remote_repo_url>nfig
git push -u <remote_name> <local_branch_name>

-start with existing project, start working on the project
git clone ssh://name@gmail.com/path/to/my-project.git
git init

-alias
alias <name>="<code>"
alias graph="git log --all --decorate --oneline --graph

-Cek status
git status
git diff --staged
git --version
git help
git log --oneline
git log -3
git log -- <style>
git log --all --decorate --oneline --graph

-Menyimpan GIT
git add <file>
git add .
git commit -m "Test"
git commit -am "Test" (file harus keadaan modified) 

-Menghapus GIT
git rm --cached
git restore <file>
git branch -D <branch> (hanya bisa klo branch sudah dimerge)
git branch -d <branch>

-Mengetahui branch yg sudah di merge
git branch --merged

-Membuat branch
git branch <branch>

-Melihat branch
git branch
git branch --list

-Pindah branch
git checkout <branch>

-Mengembalikan file yg sudah di hapus
git checkout 7667a -- <file>

-Cara pindah ke commit lain
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

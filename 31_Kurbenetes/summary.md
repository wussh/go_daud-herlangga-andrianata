Kubernetes adalah platform open source untuk mengelola kumpulan kontainer dalam suatu cluster server. Platform ini pertama kali dikembangkan oleh Google dan kini dikelola oleh Cloud Native Computing Foundation (CNCF) sebagai platform manajemen kontainer yang cukup populer.

Kontainer sendiri adalah environment dengan sumber daya, CPU, dan sistem file untuk satu aplikasi. Jadi, aplikasi tersebut akan memiliki sumber daya sendiri. Keuntungannya, aplikasi jadi tidak mudah mengalami downtime.

Kubernetes memiliki kemampuan untuk melakukan penjadwalan aplikasi, load balancing server dan peningkatan kapasitas kontainer secara otomatis. 

Komponen Kubernetes

Cluster 
Cluster adalah suatu kelompok berisi server fisik atau VPS untuk menjalankan Kubernetes. Ada dua jenis server yang dibutuhkan, yaitu master node dan worker node. 

Object
Pod merupakan objek terkecil di dalam cluster kubernetes yang terletak di dalam node. Fungsinya untuk menjalankan docker images yang membentuk sebuah kontainer. 

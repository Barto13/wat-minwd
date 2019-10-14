Przedmiot: Metody i narzędzia wspomagania decyzji

Autor: Bartłomiej Urło I6E2S1

Pakiet clp pozwala na użycie biblioteki CLP, która oryginalnie jest napisana w języku C++, używając języka GO

Biblioteka jest zaprojektowana w celu rozwiązywania matematycznych problemów optymalizacji w formie :
"Zminimalizuj c'x ,takie że: lhs ≤ Ax ≤ rhs i lb ≤ x ≤ ub"


Instalacja Go:

https://golang.org/doc/install?download=go1.13.1.windows-amd64.msi

Do uruchomienia przykładowego programu potrzeba COIN-OR Linear Programming (CLP)

Instrukcje do instalacji:

You can obtain the source code for the Clp package in two ways:

1. Obtain the source directly from the COIN-OR subversion repository
   (recommended).  For this you needs the program 'svn' installed on
   your machine, and output of "svn --version" must contain 
   "handles 'https' scheme".

   Assuming that you want to download the code into a subdirectory
   "COIN-Clp", you type

   svn co https://projects.coin-or.org/svn/Clp/trunk Coin-Clp

2. Download the tarball from http://www.coin-or.org/Tarballs/Clp and
   extract it, for example, with

   gunzip Clp_2006Jun07.tgz
   tar xvf Clp_2006Jun07.tar

   (Here "2006Jun07" is of course replaced by the string in the
   tarball you downloaded.)

   More detailed download instructions can be found at

   https://projects.coin-or.org/BuildTools/wiki/user-download
   
   
   
   
   Polecenie do potwierdzenia zainstalowania CLP 
   
   pkg-config --libs clp
   
   Następnie trzeba zainstalować bibliotekę clp za pomocą:
   
   go get github.com/lanl/clp
   
   
   
   
   

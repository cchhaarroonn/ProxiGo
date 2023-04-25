ProxiGo
ProxiGo is a command-line tool written in Golang that helps you gather proxies from websites and save them into a text file.

Installation
Before installing ProxiGo, ensure that Golang is installed on your system. Then, follow these steps:

Clone the repository:
bash
Copy code
git clone https://github.com/cchhaarroonn/ProxiGo.git
Navigate to the project directory:
bash
Copy code
cd proxigo
Build the project:
go
Copy code
go build
Run ProxiGo:
bash
Copy code
./proxigo


Usage
ProxiGo requires a list of websites from which it will gather proxies. You can provide a list of websites in a text file, with one website per line.

For example, create a file called websites.txt with the following content:

arduino
Copy code
https://www.sslproxies.org/
https://free-proxy-list.net/
https://www.proxy-list.download/
Then, run ProxiGo with the following command:

lua
Copy code
./proxigo -input websites.txt -output proxies.txt
This command will gather proxies from the websites in websites.txt and save them into a file called proxies.txt.

Options
ProxiGo supports the following command-line options:

vbnet
Copy code
-input string
    Path to the file containing a list of websites to scrape (default "websites.txt")
-output string
    Path to the output file (default "proxies.txt")
-timeout int
    Timeout for each request in seconds (default 10)
Contributing
Contributions are welcome! Please create a pull request with your changes.

License
ProxiGo is licensed under the MIT license. See LICENSE for more details.

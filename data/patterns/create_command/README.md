# Create Command

During penetration tests, many different tools are used, and often they are run with different parameters and switches depending on the target and circumstances. Because there are so many tools, it's easy to forget how to run certain tools, and what the different parameters and switches are. Most tools include a "-h" help switch to give you these details, but it's much nicer to have AI figure out all the right switches with you just providing a brief description of your objective with the tool.

# Requirements

You must have the desired tool installed locally that you want Fabric to generate the command for. For the examples above, the tool must also have help documentation at "tool -h", which is the case for most tools.

# Examples

For example, here is how it can be used to generate different commands

## sqlmap

**prompt**
```text
tool=sqlmap;echo -e "use $tool target https://example.com?test=id url, specifically the test parameter. use a random user agent and do the scan aggressively with the highest risk and level\n\n$($tool -h 2>&1)" | fabric --pattern create_command
```

**result**

```text
python3 sqlmap -u https://example.com?test=id --random-agent --level=5 --risk=3 -p test
```

## nmap

**prompt**

```text
tool=nmap;echo -e "use $tool to target all hosts in the host.lst file even if they don't respond to pings. scan the top 10000 ports and save the output to a text file and an xml file\n\n$($tool -h 2>&1)" | fabric --pattern create_command
```

**result**

```text
nmap -iL host.lst -Pn --top-ports 10000 -oN output.txt -oX output.xml
```

## gobuster

**prompt**
```text
tool=gobuster;echo -e "use $tool to target example.com for subdomain enumeration and use a wordlist called big.txt\n\n$($tool -h 2>&1)" | fabric --pattern create_command
```
**result**

```text
gobuster dns -u example.com -w big.txt
```

## dirsearch

**prompt**

```text
tool=dirsearch;echo -e "use $tool to enumerate https://example.com. ignore 401 and 404 status codes. perform the enumeration recursively and crawl the website. use 50 threads\n\n$($tool -h 2>&1)" | fabric --pattern create_command
```

**result**

```text
dirsearch -u https://example.com -x 401,404 -r --crawl -t 50
```

## nuclei

**prompt**
```text
tool=nuclei;echo -e "use $tool to scan https://example.com. use a max of 10 threads. output result to a json file. rate limit to 50 requests per second\n\n$($tool -h 2>&1)" | fabric --pattern create_command
```
**result**
```text
nuclei -u https://example.com -c 10 -o output.json -rl 50 -j
```

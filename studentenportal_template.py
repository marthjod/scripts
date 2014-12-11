#!/usr/bin/python
# -*- coding: latin-1 -*-

# Fetch grade overview from uni's portal 
# and parse for current semester's exam pass/fail 
# using Python Selenium Webdriver API

# Use as cronjob, f.ex.
#
# python studentenportal.py > /tmp/grades.txt
# scp /tmp/grades.txt login@uni:
# ssh login@uni "cat grades.txt | mailx -s Grades you@example.com"


from selenium import webdriver
import re
# against code-prying eyes (not hands!)
import base64

def main():
    driver = webdriver.Firefox()
    driver.implicitly_wait(30)
    base_url = "https://hisqis.<your institution here>/qisserver/rds?state=user&type=0"
    driver.get(base_url + "/qisserver/rds?state=user&type=1&category=auth.login&startpage=portal.vm&breadCrumbSource=portal")
    driver.find_element_by_id("asdf").clear()
    driver.find_element_by_id("asdf").send_keys("<your login name here>")
    driver.find_element_by_id("fdsa").clear()
    driver.find_element_by_id("fdsa").send_keys(base64.b64decode("<your base64-encoded pass here>"))
    driver.find_element_by_name("submit").click()
    driver.find_element_by_link_text("Prüfungsverwaltung").click()
    driver.find_element_by_link_text("Notenspiegel").click()
    table = driver.find_element_by_xpath('//*[@id="wrapper"]/div[6]/div[2]/table[2]')
    
    for line in table.text.split("\n"):
        found = re.search("<current semester>", line)
        if found:
            ignore = re.search("<lines containing this text will be ignored>", line)
            if ignore is None: 
                passed = re.search("bestanden", line)
                if passed:
                    failed = re.search("nicht bestanden", line)
                    if failed:
                        print "[!] " + line.encode('utf-8')                
                    else:                 
                        print "[x] " + line.encode('utf-8')    
                else:
                    print "[ ] " + line.encode('utf-8')

    # log out
    driver.find_element_by_link_text("Abmelden").click()
    driver.quit()
            
if __name__ == "__main__":
    main()    

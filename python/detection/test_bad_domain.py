from dataclasses import dataclass
import bad_domain
import logging
import unittest

class TestBadDomain(unittest.TestCase):
    def test_fmt_regex(self):
        logging.disable(logging.WARNING)
        
        @dataclass
        class TestCase:
            desc: str
            iocs: list
            want: str
        
        tests = [
            TestCase(
                desc = "Successful regex formatting.",
                iocs = [
                    "example.com",
			        "not.a.domain.example",
			        "test.google.com",
                ],
                want = ".*(example.com|not.a.domain.example|test.google.com)$"
            )
        ]
        for test in tests:
            bdr = bad_domain.BadDomainDetection(indicators=test.iocs)
            bdr.fmtRegex()
            self.assertEqual(bdr.regStr, test.want)

    def test_run(self):
        # <TODO: Implement me!>
        # If you break up your code into subfunctions (e.g. one for filter(), or for aggregate()),
        # feel free to unit test your subfunctions, rather than run(). 
        # E.g., for filter() and aggregate(), you could have:
        # def test_filter(self):
        #     [...]
        # 
        # def test_aggregate(self):
        #     [...]
        # Hint #1: helpers.py contains some proto and list[proto] comparison helper functions you can pass to self.assertTrue(...).
        # Hint #2: Testing large functions can be hard - consider breaking your code into subfunctions to make unit testing easier.
        # Hint #3: Use test_fmt_regex() above as an example of how to set up a table driven test (e.g. use TestCase for example), as well as the classes you might need!

if __name__ == "__main__":
    unittest.main()
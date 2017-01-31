# Fork intentions

Using VK API is a tricky thing, there is a lot of cases when having an error returned does not mean the event hasnt happened,
(like deleting SOME of messages, or getting ACCESS DENIED on accessing posts of deleted account).

My fork attempts to deal with such problems in a scope relevant to my projects.

What is out of scope here, e.g. the things that should be a part of API but are not put into it:

 - throttling. I keep using GetClient() wrapper around this api, which provides throttling/locking.
 - complex message decoding. Due to ever-changing and recursive nature of VK's data (forwards, attachments, different permissions, access_key's) it takes a lot more work.
 - client configuration. It's hacked to use HTTP and report less failures.

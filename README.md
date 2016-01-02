# ip2location-acl

IP2Location now provide feeds which look like this:

```
acl "US" {
	192.168.0.0/16;
	[...]
};

[...]
```

This is useful to us because we can stick them in a radix tree, enabling really fast (microseconds) lookups.

What this does is basically combine a bind-style configuration parser with a radix library for reasonably fast loading
and fast lookups.

It is otherwise boring on it's own.  You need the paid IP2Location subscription to get the ACL config files too.

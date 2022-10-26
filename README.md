# Tech Blogs
A few months back, I really got into reading tech blogs.
So I thought, wouldn't it be nice to have a script that gives you the latest blog posts of big tech companies? Of course, some existing websites let you pick specific blogs and have the entries shown to you in a dashboard.
However, the selfish bastard that I am, I wanted something of my own, hence, this!

This project is written using Golang. I use RSS feeds of tech blogs but sometimes some scraping is done on the side to get images.

![](assets/tb.jpg)
## List of Blogs
Currently, the following tech blogs are included:

* [Github](https://github.blog/category/engineering/)
* [LinkedIn](https://engineering.linkedin.com/)
* [Facebook/Meta](https://engineering.fb.com/)

Of course, merge requests are highly appreciated since I will spend little time extending this list.

## Schema
Currently, the entries have these fields:
* Blog
* Title
* Link
* Time
* Tags
* Summary
* Images

This project was written for deriving data to be served in a scenario where the entries are listed with a summary and an image.
So the images are all downloaded locally and their paths are stored in the `Images` field.
Also, a `body` field **IS NOT PRESENT** since I intended the actual reading to happen in the original source. However, this is not written in stone and I might change it in the future.

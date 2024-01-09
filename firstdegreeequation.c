







/**
 * main - Entry point
 * @argc: int
 * @argv: char array
 * Return: 0 on success, 1 if error
 */
int main(int argc, char *argv[])
{
	int fd, fd2, r, w, c;
	char buf[SIZE];

	if (argc != 3)
	dprintf(STDERR_FILENO, "Usage: cp file_from file_to\n"), exit(97);
	fd = open(argv[1], O_RDONLY);
	if (fd == -1)
	dprintf(STDERR_FILENO,
	"Error: Can't read from file %s\n", argv[1]), exit(98);
	fd2 = open(argv[2], O_CREAT | O_TRUNC | O_WRONLY, 0664);
	if (fd2 == -1)
	dprintf(STDERR_FILENO,
	"Error: Can't write to %s\n", argv[2]), exit(99);
	while ((r = read(fd, buf, SIZE)) > 0)
	{
	w = write(fd2, buf, r);
	if (w == -1)
	dprintf(STDERR_FILENO,
	"Error: Can't write to %s\n", argv[2]), exit(99);
	}
	if (r == -1)
	dprintf(STDERR_FILENO,
	"Error: Can't read from file %s\n", argv[1]), exit(98);
	c = close(fd);
	if (c == -1)
	dprintf(STDERR_FILENO,
	"Error: Can't close fd %d\n", fd), exit(100);
	c = close(fd2);
	if (c == -1)
	dprintf(STDERR_FILENO,
	"Error: Can't close fd %d\n", fd2), exit(100);
	return (0);
}




"""@author: github.com/isacarcanjo"""




#include <stdio.h>
#include <stdlib.h>
#include <getopt.h>
#include <unistd.h>
#include <assert.h>
#include <errno.h>
#include <string.h>
#include <fcntl.h>
#include <sys/stat.h>
#include <sys/types.h>
#include <tiffio.h>
#include <dlfcn.h>


typedef int temp_file_handle;
const char* program_name;
extern char** environ;
void testArg (int argc, char* argv[]) {
    printf ("the name of the programe is '%s'\n", argv[0]);
    printf ("the programe is invoked with %d arguments.\n", argc - 1);
    if (argc > 1) {
        int i;
        printf ("the arguments are :\n");
        for (i = 1; i < argc; ++i) {
            printf ("%s\n", argv[i]);
        }
    }
}

void print_usage (FILE * stream, int exit_code) {
    fprintf (stream, "Usage: %s options [ inputfile ... ]\n", program_name );
    fprintf (stream, 
        " -h --help             display this usage information.\n"
        " -o --output filename  writeoutput to a file.\n"
        " -v --verbose          print verbose message.\n");
    //exit (exit_code);
}

void testGetOpt (int argc, char* argv[]) {
    int nextOpt;
    const char* const short_options  = "ho:v";
    const struct option long_options[] = {
        { "help", 0, NULL, 'h' },
        { "output", 1, NULL, 'o' },
        { "verbose", 0, NULL, 'v' },
        { NULL, 0, NULL, 0 } 
    };
    const char* output_filename = NULL;
    int verbose = 0;

    program_name = argv[0];
    do {
        int next_option;
        next_option = getopt_long (argc, argv, short_options, long_options, NULL);
        switch (next_option) {
            case 'h': /* -h or --help */
                /* User has requested usage information. Print it to standard
                output, and exit with exit code zero (normal termination). */
                print_usage (stdout, 0);
            case 'o': /* -o or --output */
                /* This option takes an argument, the name of the output file. */
                output_filename = optarg;
                break;
            case 'v': /* -v or --verbose */
                verbose = 1;
                break;
            case '?': /* The user specified an invalid option. */
                /* Print usage information to standard error, and exit with exit
                code one (indicating abnormal termination). */
                print_usage (stderr, 1);
            case -1: /* Done with options. */
                break;
            default: /* Something else: unexpected. */
                abort ();
        }
    }while (nextOpt != -1);

    if (verbose) {
        int i;
        for (i = optind; i < argc; i++) {
            printf ("Argument: %s\n", argv[i]);
        }
    }
}

void testStdBuffer () {
    int i = 0;
    do{
        i++;
        printf (".");
        //sleep (1);
    }while (i < 100);
    printf ("\n");
    i = 0;
    do{
        i++;
        printf (".");
        //sleep (1);
    }while (i < 100);
    fflush (stdout);
    printf ("\n");
}


// read env
void testEnv (char* name) {
    char** var;
    for (var = environ; *var != NULL; ++var) {
        //printf ("%s\n", *var);
    }

    char* server_name = getenv ("SERVER_NAME");
    if (server_name == NULL)
    {
        server_name = "test.com";
        setenv ("SERVER_NAME", server_name, 1);
        printf ("env:SERVER_NAME=%s\n", getenv ("SERVER_NAME"));
    }
    unsetenv ("SERVER_NAME");
}

temp_file_handle testMkstemp () {
    char* buffer;
    buffer = "hello";
    int length = sizeof (buffer);
    char temp_filename[] = "/tmp/temp_filename.XXXXXX";
    int fd = mkstemp (temp_filename);
    unlink (temp_filename);
    write (fd, buffer, length);
    return fd;
}

char* read_temp_file (temp_file_handle temp_file, size_t* length) {
    //assert (length == 0);
    char* buffer;
    int fd = temp_file;
    lseek (fd, 0, SEEK_SET);
    //printf ("%d\n", (int)*length);
    //read (fd, length, sizeof (*length));
    buffer = (char*) malloc (*length);
    //printf ("%d\n", (int)*length);
    read (fd, buffer, *length);
    close (fd);
    return buffer;
}

void testErrno () {
    int fd = open ("abc.txt", O_RDONLY);
    if (fd != 0) {
        fprintf (stderr, "error opening file: %s\n", strerror (errno));
    }

    char* path = "main.c";
    int rval = chown ("main.c", 500, -1);
    assert (rval == -1);
    if (rval != 0) {
        fprintf (stderr, "error opening file: %s\n", strerror (errno));
        int error_code = errno;
        switch (error_code) {
            case EPERM:
            case EROFS:
            case ENAMETOOLONG: case ENOENT:
            case ENOTDIR: 
            case EACCES:
                fprintf (stderr, "error changing ownership of %s: %s\n", path, strerror (error_code));
            break;
            case EFAULT:
                abort ();
        }
    }else{
        fprintf (stdout, "chown success\n");
    }
}

//typedef int ssize_t;
char* testFreeResource (const char* filename, size_t length) {
    char* buffer;
    int fd;
    ssize_t bytes_read;
    buffer = (char*) malloc (length);
    if (buffer == NULL)
        return NULL;
    fd = open (filename, O_RDONLY);
    if (fd == -1) {
        free(buffer);
        return NULL;
    }

    bytes_read = read (fd, buffer, length);
    if (bytes_read != length) {
        free(buffer);
        close(fd);
        return NULL;
    }
    close(fd);
    fprintf(stdout, "buffer: %s\n", buffer);
    return buffer;
}

void testLibDependency(){
    TIFF* tiff;
    tiff = TIFFOpen("", "r");
    TIFFClose(tiff);
}

void testDynamicLoading(){
    void* handel = dlopen("libtest.so" ,RTLD_LAZY);
    int (*test)() = dlsym(handel, "f");
    if (*test == NULL){
        char * err = dlerror();
        printf("%s", err);
    } 
    int a = (*test)();
    dlclose(handel);
    printf("dlopen  dlsym result: %d\n", a);
}

//argc argv
int main (int argc, char* argv[]) {
    testDynamicLoading();
    testFreeResource ("main.c", 5);
    testErrno ();
    testEnv ("PATH");
    testStdBuffer ();
    temp_file_handle fd = testMkstemp ();
    size_t length = 5;
    char* buffer = read_temp_file (fd, &length);
    printf ("read buffer: %s\n", buffer);
    testArg (argc, argv);
    testGetOpt (argc, argv);
}


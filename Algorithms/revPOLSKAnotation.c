#include <stdio.h>
#include <stdlib.h>

typedef struct s_vals {
    int num;
    struct s_vals *next;
}         t_vals;

typedef struct s_zn {
    char znak;
    struct s_zn *next;
}               t_zn;

int is_znak(char c) {
    if (c == '+' || c == '-' || c == '*' || c == '/' || c == '%') {
        return 0;
    }
    return 1;
}

int is_num(char c) {
    if (c < '0' || c > '9') {
        return 1;
    }
    return 0;
}

int is_neg_num(char *num) {
    int i = 0;
    if (num[0] == '-') {
        i++;
    }
    while (num[i] && num[i] != ' ') {
        if (is_num(num[i++])) {
            return 1;
        }
    }
    return 0;
}

int ft_strlen(char *av) {
    int i = 0;
    while (av[i]) {
        i++;
    }
    return i;
}

int mini_atoi(char *str) {
    int i = 0, res = 0;
    int neg = 0;
    if (str[i] == '-') {
        neg = 1;
        i++;
    }
    while (str[i]) {
        res = res * 10 + str[i++] - '0';
    }
    if (neg == 1) {
        return -res;
    }
    return res;
}

int push_num(t_vals **vals, int num) {
    if (*vals == 0) {
        if (!(*vals = (t_vals*)malloc(sizeof(t_vals)))) {
            return 1;
        }
        (*vals)->num = num;
        (*vals)->next = 0;
    } else {
        t_vals *tmp = *vals;
        t_vals *new = 0;
        if (!(new = (t_vals*)malloc(sizeof(t_vals)))) {
            return 1;
        }
        new->num = num;
        new->next = tmp;
        *vals = new;
    }
    return 0;
}

void pop_num(t_vals **vals) {
    t_vals *tmp = *vals;
    *vals = (*vals)->next;
    free(tmp);
}

int calculate(t_vals **vals, char zn, char fl) {
    t_vals *start = *vals;

    if (!start) {
        return 1;
    } else if (!(start->next) && fl != 'l') {
        return 1;
    } else if (start->next && fl == 'l') {
        return 1;
    } else if (!(start->next) && fl == 'l') {
        printf("%d\n", start->num);
        return 0;
    }
    int r = start->num, l = start->next->num;
    int res = 0;
    switch (zn) {
        case '+': {
            res = l + r;
            break;
        }
        case '-': {
            res = l - r;
            break;
        }
        case '*': {
            res = l * r;
            break;
        }
        case '/': { // if r == 0, send err
            res = l / r;
            break;
        }
        default: { // '%'
            res = l % r;
        }
    }
    pop_num(vals);
    pop_num(vals);
    push_num(vals, res);
    return 0;
}

int rpn(char *av) {
    int i = 0;
    int len = ft_strlen(av);
    if (is_neg_num(av) || is_znak(av[len-1])) {
        return 1;
    }
    int l = 0;
    t_vals *nums = 0;
    while (av[i]) {
        if (av[i] != ' ' && is_znak(av[i]) && is_num(av[i])) {
            return 1;
        } else if (l == -1 && av[i] != ' ') {
            l = i;
        }
        if (l != -1 && (av[i + 1] == ' ' || av[i + 1] == '\0')) {
            if (i == l && !is_znak(av[i])) {
                if (calculate(&nums, av[i], '1')) {
                    return 1;
                }
            } else {
                char *tmp = 0;
                tmp = (char *) malloc(sizeof(char) * (i - l + 2));
                if (!tmp) {
                    return 1;
                }
                int j = 0;
                while (l <= i) {
                    tmp[j++] = av[l++];
                }
                tmp[j] = '\0';
                if (!is_neg_num(tmp)) {
                    int num = mini_atoi(tmp);
                    if (push_num(&nums, num)) {
                        return 1;
                    }
                } else {
                    return 1;
                }
                free(tmp);
            }
            l = -1;
        }
        i++;
    }
    if (!nums) {
        return 1;
    }
    return (calculate(&nums, '-', 'l'));
}

int main(int ac, char **av) {
    if (ac == 2 && *av[1]) {
        if (rpn(av[1])) {
            printf("Error\n");
        }
    } else {
        printf("Error\n");
    }
}
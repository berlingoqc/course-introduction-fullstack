---
title: "Setuping development environment"
date: 2022-05-27T17:08:19-03:00
weight: 30
---

Now it's time to work, we are gonna setup you with a development environment.
To write code we need many tools to help us, they are all part of our development
envrionment.

It contains the following component normally:

* A text editor
* A compiler or interpreter for your language
* Tooling to help you

In a further chapiter we will talk about IDE (Integrated Development Environment) that 
bundle all of this together but for your own good we will learn to make our own IDE with
great individual software, this will gave you more flexibility and understanding in the
long run. IDE are really usefull but me personnally i only use them when i have special
need, like integration with tooling. We will use instead a source code editor.

For me there is only two option for this.

* Visual Studio Code (with the vim plugin)
* Vim (nvim)


I'll introduce both of them to you here and you can make your own decision. But
one thing for sure you have to learn the basic of Vim.

## Introduction to Vim

Vim is a terminal based editor that exists since the time when mouse where not really 
a thing. It has mouse support but not by default , every thing that you what to do
you must do it with shortcut. It can be very painfull at first but the speed, efficency
, portability and lightness will be a huge benifice for you in the long run.

Here is a guide to setup you up in macos with nvim a fork of vim that has better plugin
support and modern feature.


## Before we can install , i must deviate of package manager

If you use linux you are already familiar with it but for macos and windows people it will
be a new feature for you. A link is provide bellow to learn more about this but fastly it's

> A package manager or package-management system is a collection of software tools that automates the process of installing, upgrading, configuring, and removing computer programs for a computer in a consistent manner

With this you can easily with command install all the software that you need. You will see soon enought why it's a huge deal.

Macos does not come with a package manager but there is [brew](https://brew.sh/). We will install it and use it to setup your development environment.

```bash
$ /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

Tadam you now have brew install so we can move to configure nvim and golang afterward.

## Installing and configuration NVIM

```bash
$ brew install neovim
```

After we will create the following configuration file (this is my configuration file without what you don't need)

```bash
# Create this file with the following content
➜  golang_hello_world git:(master) ✗ cat $HOME/.config/nvim/init.vim
set runtimepath^=~/.vim runtimepath+=~/.vim/after
let &packpath = &runtimepath
source ~/.vimrc


# Use this command to download VimPlug a plugin manager for vim to add extra feature
# https://github.com/junegunn/vim-plug
➜  golang_hello_world git:(master) ✗ curl -fLo ~/.vim/autoload/plug.vim --create-dirs \
    https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim
```

```vim
" $HOME/.vimrc
syntax enable
set tabstop=4
set softtabstop=4
set expandtab
set number
set cursorline
set noswapfile

filetype indent on
filetype plugin on
filetype plugin indent on

set autowrite
set mouse=a

nnoremap <c-z> <nop>

map <C-n> :cnext<CR>
map <C-m> :cprevious<CR>
nnoremap <leader>a :cclose<CR>

set laststatus=2

call plug#begin('~/.vim/plugged')

Plug 'scrooloose/nerdtree', { 'on':  'NERDTreeToggle' }
Plug 'fatih/vim-go', { 'tag': '*' }
Plug 'altercation/vim-colors-solarized'
Plug 'vim-airline/vim-airline'
Plug 'vim-airline/vim-airline-themes'
Plug 'https://github.com/mcchrish/zenbones.nvim'
Plug 'https://github.com/leafgarland/typescript-vim'
Plug 'https://github.com/Quramy/vim-js-pretty-template'
Plug 'https://github.com/Shougo/vimproc.vim', {'do' : 'make'}
Plug 'https://github.com/vim-syntastic/syntastic'
Plug 'https://github.com/editorconfig/editorconfig-vim'
Plug 'https://github.com/leafgarland/typescript-vim'
Plug 'neoclide/coc.nvim', {'branch': 'release'}
Plug 'https://github.com/ctrlpvim/ctrlp.vim'

Plug 'mdempsky/gocode', { 'rtp': 'vim', 'do': '~/.vim/plugged/gocode/vim/symlink.sh' }
Plug 'fatih/vim-go', { 'do': ':GoUpdateBinaries' }
Plug 'udalov/kotlin-vim'
Plug 'https://github.com/editorconfig/editorconfig-vim'
Plug 'morhetz/gruvbox'
Plug 'lifepillar/vim-solarized8'
Plug 'airblade/vim-gitgutter'
Plug 'tpope/vim-fugitive'
Plug 'Xuyuanp/nerdtree-git-plugin'
Plug 'dracula/vim', { 'name': 'dracula' }
Plug 'christoomey/vim-tmux-navigator'


Plug 'alvan/vim-closetag'
  let g:closetag_filenames = '*.html,*.xhtml,*.xml,*.vue,*.phtml,*.js,*.jsx,*.coffee,*.erb'



call plug#end()

set termguicolors
set background=light " or dark

colorscheme forestbones

:nnoremap <C-g> :NERDTreeToggle<CR>

set wildignore+=*/dist/*,*/node_modules/*
let g:ctrlp_working_path_mode = 'ra'

let g:ctrlp_map = '<c-p>'
let g:ctrlp_cmd = 'CtrlP'

let g:ycm_rust_src_path = '/opt/rust/src'

let g:typescript_compiler_binary = 'tsc'
let g:typescript_compiler_options = ''

autocmd QuickFixCmdPost [^l]* nested cwindow
autocmd QuickFixCmdPost l* nested cwindow

autocmd FileType typescript syn clear foldBraces

set statusline+=%#warningmsg#
set statusline+=%{SyntasticStatuslineFlag()}
set statusline+=%*
let g:syntastic_check_on_open = 1
let g:syntastic_check_on_wq = 0
let g:tsuquyomi_disable_quickfix = 1
let g:syntastic_typescript_checkers = ['tsuquyomi']

if !exists('g:syntastic_html_tidy_ignore_errors')
  let g:syntastic_html_tidy_ignore_errors = []
endif
let g:syntastic_html_tidy_ignore_errors += [ "<alb-", "discarding unexpected </alb-"]
let g:syntastic_html_tidy_ignore_errors += [ "<ng-", "discarding unexpected </ng-"]
let g:syntastic_html_tidy_ignore_errors += [ "<mat-", "discarding unexpected </mat-"]
let g:syntastic_html_tidy_ignore_errors += [ "<app-", "discarding unexpected </app-"]


let g:go_fmt_command = "goimports"
let g:go_def_mode = 'gopls'
let g:go_info_mode = 'gopls'
let g:coc_disable_startup_warning = 1




" -------------------------------------------------------------------------------------------------
"  " coc.nvim default settings
"  "
"  -------------------------------------------------------------------------------------------------

" if hidden is not set, TextEdit might fail.
set hidden
"  " Better display for messages
set cmdheight=2
"  " Smaller updatetime for CursorHold & CursorHoldI
set updatetime=300
"  " don't give |ins-completion-menu| messages.
set shortmess+=c
"  " always show signcolumns
set signcolumn=yes
"
"  " Use tab for trigger completion with characters ahead and navigate.
"  " Use command ':verbose imap <tab>' to make sure tab is not mapped by other
"  plugin.
inoremap <silent><expr> <TAB>
    \ pumvisible() ? "\<C-n>" :
    \ <SID>check_back_space() ? "\<TAB>" :
    \ coc#refresh()
inoremap <expr><S-TAB> pumvisible() ? "\<C-p>" : "\<C-h>"

function! s:check_back_space() abort
    let col = col('.') - 1
    return !col || getline('.')[col - 1]  =~# '\s'
endfunction

"                        " Use <c-space> to trigger completion.
inoremap <silent><expr> <c-space> coc#refresh()
"
"                        " Use `[c` and `]c` to navigate diagnostics
nmap <silent> [c <Plug>(coc-diagnostic-prev)
nmap <silent> ]c <Plug>(coc-diagnostic-next)
"
"                        " Remap keys for gotos
nmap <silent> gd <Plug>(coc-definition)
nmap <silent> gy <Plug>(coc-type-definition)
nmap <silent> gi <Plug>(coc-implementation)
nmap <silent> gr <Plug>(coc-references)
"
"                        " Use U to show documentation in preview window
nnoremap <silent> U :call <SID>show_documentation()<CR>
"
"                        " Remap for rename current word
nmap <leader>rn <Plug>(coc-rename)
"
"                        " Remap for format selected region
vmap <leader>f  <Plug>(coc-format-selected)
nmap <leader>f  <Plug>(coc-format-selected)
"                        " Show all diagnostics
nnoremap <silent> <space>a  :<C-u>CocList diagnostics<cr>
"                        " Manage extensions
nnoremap <silent> <space>e  :<C-u>CocList extensions<cr>
"                        " Show commands
nnoremap <silent> <space>c  :<C-u>CocList commands<cr>
"                        " Find symbol of current document
nnoremap <silent> <space>o  :<C-u>CocList outline<cr>
"                        " Search workspace symbols
nnoremap <silent> <space>s  :<C-u>CocList -I symbols<cr>
"                        " Do default action for next item.
nnoremap <silent> <space>j  :<C-u>CocNext<CR>

"                        " Do default action for previous item.
nnoremap <silent> <space>k  :<C-u>CocPrev<CR> "                        " Resume latest coc list
nnoremap <silent> <space>p  :<C-u>CocListResume<CR>
```


Open nvim and enter the following command to install the plugin `:PlugInstall`

You should now be all setup to start coding with vim.


## Install on configure go

```bash
$ brew install go
# You should now be have the go binary
# You can try it and see the version
➜  golang_hello_world git:(master) ✗ go version
go version go1.18.1 darwin/amd64
```


## Exercice 

* vitutor
* Do the hello world program with vim and go

## Things to read

* https://www.tutorialspoint.com/vim/index.htm
* https://en.wikipedia.org/wiki/Package_manager
* https://octetz.com/docs/2019/2019-04-24-vim-as-a-go-ide/









import { Button } from "@/components/ui/button"
import {useEffect, useState} from "react";



type PaginationNo = (number | '...')[]
type PaginationNoProps = {
  total: number,
  pageNum: number,
  pageSize: number
}

const getPaginationNo = (p:PaginationNoProps): PaginationNo => {
  const totalPage = Math.ceil(p.total / p.pageSize);
  const pages: PaginationNo = [];

  if (totalPage <= 9) {
    for (let i = 1; i <= totalPage; i++) {
      pages.push(i);
    }
  } else {
    if (p.pageNum <= 6) {
      for (let i = 1; i <= 9; i++) {
        pages.push(i);
      }
      pages.push('...');
      pages.push(totalPage);
    } else if (p.pageNum > 6 && p.pageNum <= (totalPage - 4)) {
      pages.push(1);
      pages.push('...');
      for (let i = (p.pageNum - 4); i <= (p.pageNum + 4); i++) {
        pages.push(i);
      }
      pages.push('...');
      pages.push(totalPage);
    } else {
      pages.push(1);
      pages.push('...');
      for (let i = (totalPage - 8); i <= totalPage; i++) {
        pages.push(i);
      }
    }
  }

  return pages;
}

export const PaginationComponent = (p: PaginationNoProps & {onPageChange: (page: number) => void}) => {

  const [paginationNo, setPaginationNo] = useState<PaginationNo>([])
  useEffect(() => {
    setPaginationNo(getPaginationNo(p))
  }, [p])

  return (
      <div className="flex items-center justify-center space-x-4">
        <Button
            disabled={p.pageNum === 1 || paginationNo.length ===0}
            onClick={() => p.onPageChange(p.pageNum - 1)} className="bg-white px-4 py-2 rounded-md text-sm font-medium text-gray-700 hover:bg-gray-200 focus:outline-none">
          <ArrowLeftIcon className="w-4 h-4" />
          Prev
        </Button>
        {
          paginationNo.map(flag => {
            if(flag==='...'){
              return <div key={flag} className="px-2 py-1 text-sm font-medium text-gray-700">...</div>
            }
            if(p.pageNum === flag){
              return <Button key={flag}  disabled className="bg-white px-4 py-2 rounded-md text-sm font-medium text-gray-700 hover:bg-gray-200 focus:outline-none">
                { flag }
              </Button>
            }
            return <Button key={flag}  className="bg-white px-4 py-2 rounded-md text-sm font-medium text-gray-700 hover:bg-gray-200 focus:outline-none" onClick={() => p.onPageChange(flag)}>
              { flag }
            </Button>
          })
        }
        <Button
            disabled={(p.pageNum === Math.ceil(p.total / p.pageSize))||paginationNo.length===0}
            onClick={() => p.onPageChange(p.pageNum+1)}
            className="bg-white px-4 py-2 rounded-md text-sm font-medium text-gray-700 hover:bg-gray-200 focus:outline-none">
          Next
          <ArrowRightIcon className="w-4 h-4" />
        </Button>
      </div>
  )
}

function ArrowLeftIcon(props:{className:string}) {
  return (
      <svg
          {...props}
          xmlns="http://www.w3.org/2000/svg"
          width="24"
          height="24"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          strokeWidth="2"
          strokeLinecap="round"
          strokeLinejoin="round"
      >
        <path d="m12 19-7-7 7-7" />
        <path d="M19 12H5" />
      </svg>
  )
}


function ArrowRightIcon(props:{className:string}) {
  return (
      <svg
          {...props}
          xmlns="http://www.w3.org/2000/svg"
          width="24"
          height="24"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          strokeWidth="2"
          strokeLinecap="round"
          strokeLinejoin="round"
      >
        <path d="M5 12h14" />
        <path d="m12 5 7 7-7 7" />
      </svg>
  )
}

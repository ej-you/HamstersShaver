import os

from pytonapi import AsyncTonapi
from pytonapi.schema.events import TransactionEventData

# Enter your API key
API_KEY = os.getenv("TON_API_KEY")

# List of TON blockchain accounts to monitor
ACCOUNTS = [
    "UQDLjMTSj58tHMc9oIeo00dbV6smPp1plQw5vUSZADeUInM9",

    "UQCJYWj8oz1sceKGTqLELUhJ9jAJynXRiRW2AqDpTb-8utwN",
    "UQDJmlxQ6Bd7N_tmHuf2p-imf8H5CwIunUSJI5C15H5gUqTY",
    "UQADRQ7WjDlbuNLw7jWKSWE7Si_oA1t_JWDyXtotfCyaP0qY",
    "UQCYbWPeHQSw4lwUBIwcnxR6avVwIpax4f-xR-ba2N6A0CEF",
    "EQCuqtljI1F4CILgj9o-th3vqB9eNF_mt3jtyDIEWz--pojh",
    "UQDEd5vftG87eBetMbGrb2BQdKgj0bgkJqMm6IymmVeIlTtM",
    "UQC9yHiF6CtrtwLpQZfBzq8YiLUrczBPq4IQt5H8tili-gbD",
    "UQDDFR4bnkDaKoFk26gaOkxfnsFlhtt4XidUm81xfieUHIpj",
    "UQAtLX1jSk459g5y7pPqw4SkEYMnDZgxd7nG-iMBUVqwaN--",
    "UQAZgA1LHMrIiYiGmaOGSujQTasommm7LFRSrThtElr9IZuZ",
    "UQDPy5vjHGOLvEBlHpiFZgbU7jHpNoLU-FNW-ZkUcLbEb0FF",
    "UQA4mfrV45OEIuTyJKDQe41FX1X0XD8IPJ9UYb7Tpu3gK6kO",  # какой-то бот, поэтому очень хорошо его отслеживать
]


async def handler(event: TransactionEventData, tonapi: AsyncTonapi) -> None:
    """
    Handle SSEvent for transactions.

    :param event: The SSEvent object containing transaction details.
    :param tonapi: An instance of AsyncTonapi for interacting with TON API.
    """
    # trace = await tonapi.traces.get_trace(event.tx_hash)
    #
    print(event)


async def main() -> None:
    tonapi = AsyncTonapi(api_key=API_KEY)
    print("tonapi:", tonapi)

    # Subscribe to transaction events for the specified accounts
    await tonapi.sse.subscribe_to_transactions(
        handler, accounts=ACCOUNTS, args=(tonapi,)
    )


if __name__ == "__main__":
    import asyncio

    asyncio.run(main())

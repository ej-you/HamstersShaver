import asyncio

from pytoniq import LiteBalancer, LiteClient
from pytoniq_core import BlockIdExt


WORKCHAIN = 0
SHARD_STR = "e000000000000000"
SEQNO = 46902564
ROOT_HASH = "021848c68faf9eb79104bb2a7ae5e55a3df39f636c0c381cf624c7655ff39155"
FILE_HASH = "382788420be83d4069f17e0506c95bc9fbc93b0218f0ebd6c4930685e23123bf"

TRANS_LT = 50640780000001
WALLET_ADDRESS = "UQDLjMTSj58tHMc9oIeo00dbV6smPp1plQw5vUSZADeUInM9"


async def action(client: LiteClient | LiteBalancer) -> None:
    shard_int = int(SHARD_STR, 16)
    if shard_int >= 2 ** 63:
        shard_int -= 2 ** 64
    print("shard_int:", shard_int)

    print("Создание класса блока...")
    block = BlockIdExt(
        workchain=WORKCHAIN,
        shard=shard_int,
        seqno=SEQNO,
        root_hash=ROOT_HASH,
        file_hash=FILE_HASH,
    )
    print("Класс блока создан!")
    print("block:", block)

    # print("Получение информации о блоке...")
    # block_info = await client.raw_get_block(block)
    # print("Информация о блоке получена...")
    # print("block_info:", block_info)

    print("Получение информации о транзакции...")
    trans_info = await client.get_one_transaction(
        address=WALLET_ADDRESS,
        lt=TRANS_LT,
        block=block,
    )
    print("Информация о транзакции получена!")
    print("trans_info:", trans_info)


async def main() -> None:
    print("Создание клиента...")
    # client = LiteBalancer.from_mainnet_config(trust_level=2)
    # await action(client)

    async with LiteClient.from_mainnet_config(ls_i=2, trust_level=2) as client:
        print("Клиент создан!")
        print("client:", client)
        await action(client)


if __name__ == '__main__':
    asyncio.run(main())
